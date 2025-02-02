/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package test_init

import (
	"testing"

	"magma/orc8r/cloud/go/orc8r"
	accessd_test_init "magma/orc8r/cloud/go/services/accessd/test_init"
	certifier_test_init "magma/orc8r/cloud/go/services/certifier/test_init"
	"magma/orc8r/cloud/go/services/configurator"
	"magma/orc8r/cloud/go/services/configurator/protos"
	"magma/orc8r/cloud/go/services/configurator/servicers"
	"magma/orc8r/cloud/go/services/configurator/storage"
	"magma/orc8r/cloud/go/sqorc"
	"magma/orc8r/cloud/go/test_utils"
)

func StartTestService(t *testing.T) {
	db, err := sqorc.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Could not initialize sqlite DB: %s", err)
	}
	idGenerator := storage.DefaultIDGenerator{}
	storageFactory := storage.NewSQLConfiguratorStorageFactory(db, &idGenerator, sqorc.GetSqlBuilder())
	err = storageFactory.InitializeServiceStorage()
	if err != nil {
		t.Fatalf("Could not initialize storage: %s", err)
	}

	accessd_test_init.StartTestService(t)
	certifier_test_init.StartTestService(t)

	srv, lis := test_utils.NewTestService(t, orc8r.ModuleName, configurator.ServiceName)
	nb, err := servicers.NewNorthboundConfiguratorServicer(storageFactory)
	if err != nil {
		t.Fatalf("Failed to create NB configurator servicer: %s", err)
	}
	protos.RegisterNorthboundConfiguratorServer(srv.GrpcServer, nb)

	sb, err := servicers.NewSouthboundConfiguratorServicer(storageFactory)
	if err != nil {
		t.Fatalf("Failed to create SB configurator servicer: %s", err)
	}
	protos.RegisterSouthboundConfiguratorServer(srv.GrpcServer, sb)

	go srv.RunTest(lis)
}
