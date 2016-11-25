#http://www.apache.org/licenses/LICENSE-2.0.txt
#
#
#Copyright 2015 Intel Corporation
#
#Licensed under the Apache License, Version 2.0 (the "License");
#you may not use this file except in compliance with the License.
#You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
#Unless required by applicable law or agreed to in writing, software
#distributed under the License is distributed on an "AS IS" BASIS,
#WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#See the License for the specific language governing permissions and
#limitations under the License.

from modules import utils
from modules.logger import log
from unittest import TextTestRunner

import sys
import unittest


class FacterCollectorLargeTest(unittest.TestCase):

    def setUp(self):
        # set and download required binaries (snapteld, snaptel, plugins)
        self.binaries = utils.set_binaries()
        utils.download_binaries(self.binaries)

        log.debug("Starting snapteld")
        self.binaries.snapteld.start()
        if not self.binaries.snapteld.isAlive():
            self.fail("snapteld thread died")

        log.debug("Waiting for snapteld to finish starting")
        if not self.binaries.snapteld.wait():
            log.error("snapteld errors: {}".format(self.binaries.snapteld.errors))
            self.binaries.snapteld.kill()
            self.fail("snapteld not ready, timeout!")

    def test_facter_collector_plugin(self):
        # load facter collector
        loaded = self.binaries.snaptel.load_plugin("snap-plugin-collector-facter")
        self.assertTrue(loaded, "facter collector loaded")

        # check available metrics, plugins and tasks
        metrics = self.binaries.snaptel.list_metrics()
        plugins = self.binaries.snaptel.list_plugins()
        tasks = self.binaries.snaptel.list_tasks()
        self.assertGreater(len(metrics), 0, "Metrics available {} expected {}".format(len(metrics), 0))
        self.assertEqual(len(plugins), 1, "Plugins available {} expected {}".format(len(plugins), 1))
        self.assertEqual(len(tasks), 0, "Tasks available {} expected {}".format(len(tasks), 0))

        # check config policy for metric
        rules = self.binaries.snaptel.metric_get("/intel/facter/id")
        self.assertEqual(len(rules), 1, "Rules available {} expected {}".format(len(rules), 1))

        # create and list available task
        task_id = self.binaries.snaptel.create_task("/snap-plugin-collector-facter/scripts/docker/large/task.json")
        tasks = self.binaries.snaptel.list_tasks()
        self.assertEqual(len(tasks), 1, "Tasks available {} expected {}".format(len(tasks), 1))

        # check if task hits and fails
        hits = self.binaries.snaptel.task_hits_count(task_id)
        fails = self.binaries.snaptel.task_fails_count(task_id)
        self.assertGreater(hits, 0, "Task hits {} expected {}".format(hits, ">0"))
        self.assertEqual(fails, 0, "Task fails {} expected {}".format(fails, 0))

        # stop task and list available tasks
        stopped = self.binaries.snaptel.stop_task(task_id)
        self.assertTrue(stopped, "Task stopped")
        tasks = self.binaries.snaptel.list_tasks()
        self.assertEqual(len(tasks), 1, "Tasks available {} expected {}".format(len(tasks), 1))

        # unload plugin, list metrics and plugins
        self.binaries.snaptel.unload_plugin("collector", "facter", "8")
        metrics = self.binaries.snaptel.list_metrics()
        plugins = self.binaries.snaptel.list_plugins()
        self.assertEqual(len(metrics), 0, "Metrics available {} expected {}".format(len(metrics), 0))
        self.assertEqual(len(plugins), 0, "Plugins available {} expected {}".format(len(plugins), 0))

        # check for snapteld errors
        self.assertEqual(len(self.binaries.snapteld.errors), 0, "Errors found during snapteld execution")

    def tearDown(self):
        log.debug("Stopping snapteld thread")
        self.binaries.snapteld.stop()
        if self.binaries.snapteld.isAlive():
            log.warn("snapteld thread did not died")

if __name__ == "__main__":
    test_suite = unittest.TestLoader().loadTestsFromTestCase(FacterCollectorLargeTest)
    test_result = TextTestRunner().run(test_suite)
    # exit with return code equal to number of failures
    sys.exit(len(test_result.failures))
