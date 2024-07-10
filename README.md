

# Protostar
Metrics aggregator for c12s platform

## API Documentation

### GET /api/metrics-api/{timestamp}/{nodeId}

#### Description
The endpoint retrieves node metrics starting from a specified timestamp.

#### Request

##### URL
`http://localhost:8086/api/metrics-api/{timestamp}/{nodeId}`

##### Method
`GET`

##### URL Path Parameters
| Parameter | Type   | Description                                                  |
|-----------|--------|--------------------------------------------------------------|
| timestamp | string | Unix value of the timestamp from which we want values to start. |
| nodeId    | string | Identification of the node from which we want to collect our metric data. |

##### Request Headers
None

##### Example
Using Postman:
```
http://localhost:8086/api/metrics-api/1718723240/b2696545-7931-47c7-af11-cc2ce1d5e10d
```

#### Response - 201 Success
```json
{
    "status": 201,
    "data": {
        "status": "success",
        "data": {
            "resultType": "matrix",
            "result": [
                {
                    "metric": {
                        "__name__": "container_cpu_usage_seconds_total",
                        "container_label_com_docker_compose_config_hash": "0564f715245d3d7a8b140ce68be65b5246116cfcc963c1408673c66469716c00",
                        "container_label_com_docker_compose_container_number": "1",
                        "container_label_com_docker_compose_depends_on": "rate_limiter_service:service_started:false,magnetar:service_started:false,oort:service_started:false,apollo:service_started:false,kuiper:service_started:false,quasar:service_started:false",
                        "container_label_com_docker_compose_image": "sha256:bf47588af42892bb70b9ca6a0d41d50928cdd5baa9f984b2bb0937f9a9efc6c4",
                        "container_label_com_docker_compose_oneoff": "False",
                        "container_label_com_docker_compose_project": "tools",
                        "container_label_com_docker_compose_project_config_files": "/home/bunjo/Desktop/college/projects/c12s/tools/docker-compose.yml",
                        "container_label_com_docker_compose_project_working_dir": "/home/bunjo/Desktop/college/projects/c12s/tools",
                        "container_label_com_docker_compose_service": "lunar-gateway",
                        "container_label_com_docker_compose_version": "2.27.0",
                        "cpu": "total",
                        "id": "/docker/913ef1bcab1e47e25130e74eb499da81e33feaba9bcb38af311a11719553dcbf",
                        "image": "lunar-gateway",
                        "instance": "health-check:8080",
                        "job": "health-check",
                        "name": "lunar-gateway",
                        "nodeID": "b2696545-7931-47c7-af11-cc2ce1d5e10d"
                    },
                    "values": [
                        [
                            1720514135,
                            "0.034371"
                        ],
                        [
                            1720514851,
                            "0.056788"
                        ]
                    ]
                },
                {
                    "metric": {
                        "__name__": "container_cpu_usage_seconds_total",
                        "container_label_com_docker_compose_config_hash": "080eb67bae2cc75484bf1d96130bf808cc64e592f60b5becd19294cacb2e01b5",
                        "container_label_com_docker_compose_container_number": "1",
                        "container_label_com_docker_compose_image": "sha256:68dd2cdef1732a37435ca3fe75cfa849fbba13f7f0e615a96a1620029a3dad0e",
                        "container_label_com_docker_compose_oneoff": "False",
                        "container_label_com_docker_compose_project": "tools",
                        "container_label_com_docker_compose_project_config_files": "/home/bunjo/Desktop/college/projects/c12s/tools/docker-compose.yml",
                        "container_label_com_docker_compose_project_working_dir": "/home/bunjo/Desktop/college/projects/c12s/tools",
                        "container_label_com_docker_compose_service": "cassandra",
                        "container_label_com_docker_compose_version": "2.27.0",
                        "container_label_org_opencontainers_image_ref_name": "ubuntu",
                        "container_label_org_opencontainers_image_version": "22.04",
                        "cpu": "total",
                        "id": "/docker/9efcae5c59f82947c83d60bfb082ccc50c3a0f2ba2830c0798535045ea20a826",
                        "image": "tools-cassandra",
                        "instance": "health-check:8080",
                        "job": "health-check",
                        "name": "cassandra",
                        "nodeID": "b2696545-7931-47c7-af11-cc2ce1d5e10d"
                    },
                    "values": [
                        [
                            1720514135,
                            "21.141999"
                        ],
                        [
                            1720514851,
                            "52.748081"
                        ]
                    ]
                },
                {
                    "metric": {
                        "__name__": "container_cpu_usage_seconds_total",
                        "container_label_com_docker_compose_config_hash": "118aa150e527155107a4a55c9d548b48ab1b3a09f2a2d48fe105a93cdf165853",
                        "container_label_com_docker_compose_container_number": "1",
                        "container_label_com_docker_compose_depends_on": "cassandra:service_healthy:false,nats:service_started:false,vault:service_started:false",
                        "container_label_com_docker_compose_image": "sha256:fad6c9e26fe8062a94515924541f15aeef7f35ed2c12f04022b6c1825f53b6b3",
                        "container_label_com_docker_compose_oneoff": "False",
                        "container_label_com_docker_compose_project": "tools",
                        "container_label_com_docker_compose_project_config_files": "/home/bunjo/Desktop/college/projects/c12s/tools/docker-compose.yml",
                        "container_label_com_docker_compose_project_working_dir": "/home/bunjo/Desktop/college/projects/c12s/tools",
                        "container_label_com_docker_compose_service": "apollo",
                        "container_label_com_docker_compose_version": "2.27.0",
                        "container_label_restartcount": "2",
                        "cpu": "total",
                        "id": "/docker/699318abdc21cd8de6c6ca2b18fa08035ede888b96e099de7ff4bc5917029238",
                        "image": "apollo",
                        "instance": "health-check:8080",
                        "job": "health-check",
                        "name": "apollo",
                        "nodeID": "b2696545-7931-47c7-af11-cc2ce1d5e10d"
                    },
                    "values": [
                        [
                            1720514135,
                            "0.03344"
                        ],
                        [
                            1720514851,
                            "0.172057"
                        ]
                    ]
                },
                {
                    "metric": {
                        "__name__": "container_cpu_usage_seconds_total",
                        "container_label_com_docker_compose_config_hash": "3f6045278664302bce3e0e9b86c7adb0ba36153e4b549326cf18946c6566d170",
                        "container_label_com_docker_compose_container_number": "1",
                        "container_label_com_docker_compose_image": "sha256:f4641284d925d63e52041f56b440f318c4fe9f5ac54fe67206b1975dbaa3316b",
                        "container_label_com_docker_compose_oneoff": "False",
                        "container_label_com_docker_compose_project": "tools",
                        "container_label_com_docker_compose_project_config_files": "/home/bunjo/Desktop/college/projects/c12s/tools/docker-compose.yml",
                        "container_label_com_docker_compose_project_working_dir": "/home/bunjo/Desktop/college/projects/c12s/tools",
                        "container_label_com_docker_compose_service": "kuiper_etcd",
                        "container_label_com_docker_compose_version": "2.27.0",
                        "container_label_com_vmware_cp_artifact_flavor": "sha256:c50c90cfd9d12b445b011e6ad529f1ad3daea45c26d20b00732fae3cd71f6a83",
                        "container_label_org_opencontainers_image_base_name": "docker.io/bitnami/minideb:bookworm",
                        "container_label_org_opencontainers_image_created": "2024-07-04T12:49:39Z",
                        "container_label_org_opencontainers_image_description": "Application packaged by Broadcom, Inc.",
                        "container_label_org_opencontainers_image_documentation": "https://github.com/bitnami/containers/tree/main/bitnami/etcd/README.md",
                        "container_label_org_opencontainers_image_licenses": "Apache-2.0",
                        "container_label_org_opencontainers_image_ref_name": "3.5.14-debian-12-r4",
                        "container_label_org_opencontainers_image_source": "https://github.com/bitnami/containers/tree/main/bitnami/etcd",
                        "container_label_org_opencontainers_image_title": "etcd",
                        "container_label_org_opencontainers_image_vendor": "Broadcom, Inc.",
                        "container_label_org_opencontainers_image_version": "3.5.14",
                        "cpu": "total",
                        "id": "/docker/a9722c0c023ca47b2e1dbaa318cce1a622cf0b596905f6f02bcf2f6b944b9155",
                        "image": "bitnami/etcd:latest",
                        "instance": "health-check:8080",
                        "job": "health-check",
                        "name": "kuiper_etcd",
                        "nodeID": "b2696545-7931-47c7-af11-cc2ce1d5e10d"
                    },
                    "values": [
                        [
                            1720514135,
                            "0.685858"
                        ],
                        [
                            1720514851,
                            "3.092391"
                        ]
                    ]
                }
            ]
        }
    }
}
```

##### Response Body Properties (201 Success)
| Property    | Type   | Description                                        |
|-------------|--------|----------------------------------------------------|
| status      | int    | HTTP status code                                   |
| data        | object | Contains the status of the request and the metrics data |
| status      | string | Status of the data retrieval (`success`)           |
| resultType  | string | Type of result (`matrix`)                          |
| result      | array  | Array of metrics and their values                  |
| metric      | object | Contains various metric labels and their values    |
| values      | array  | Array of timestamp and metric value pairs          |

#### Response - 500 Internal Server Error
Server error with error identifier where and when in service error happened.

### GET /api/metrics-api/latest-node-data/{nodeId}

#### Description
This endpoint returns the latest custom metrics for a node that are used in our CLI.

#### Request

##### URL
`http://localhost:8086/api/metrics-api/latest-node-data/{nodeId}`

##### Method
`GET`

##### URL Path Parameters
| Parameter | Type   | Description                                     |
|-----------|--------|-------------------------------------------------|
| nodeId    | string | Identification of the node that we want data from. |

##### Request Headers
None

##### Example
Using Postman:
```
http://localhost:8086/api/metrics-api/latest-node-data/b2696545-7931-47c7-af11-cc2ce1d5e10d
```

#### Response - 201 Success
```json
{
    "status": 201,
    "data": [
        {
            "metric": {
                "__name__": "custom_node_cpu_usage_percentage",
                "instance": "health-check:8080",
                "job": "health-check",
                "nodeID": "b2696545-7931-47c7-af11-cc2ce1d5e10d"
            },
            "values": [
                [
                    1720514138,
                    "5.633947530652784"
                ]
            ]
        },
        {
            "metric": {
                "__name__": "custom_node_disk_total_gb",
                "instance": "health-check:8080",
                "job": "health-check",
                "nodeID": "b2696545-7931-47c7-af11-cc2ce1d5e10d"
            },
            "values": [
                [
                    1720514138,
                    "188.08166885375977"
                ]
            ]
        },
        {
            "metric": {
                "__name__": "custom_node_disk_usage_gb",
                "instance": "health-check:8080",
                "job": "health-check",
                "nodeID": "b2696545-7931-47c7-af11-cc2ce1d5e10d"
            },
            "values": [
                [
                    1720514138,
                    "30.84386444091797"
                ]
            ]
        }
    ]
}
```

##### Response Body Properties (201 Success)
| Property | Type | Description                              |
|----------|------|------------------------------------------|
| status   | int  | HTTP status code                         |
| data     | array| Array of metric objects and their values |

#### Response - 500 Internal Server Error
Server error with error identifier where and when in service error happened.

### GET /api/metrics-api/latest-data/{nodeId}

#### Description
This endpoint returns only the latest metrics written for the specified node.

#### Request

##### URL
`http://localhost:8086/api/metrics-api/latest-data/{nodeId}`

##### Method
`GET`

##### URL Path Parameters
| Parameter | Type   | Description                                     |
|-----------|--------|-------------------------------------------------|
| nodeId    | string | Identification of the node that we want data from. |

##### Request Headers
None

##### Example
Using Postman:
```
http://localhost:8086/api/metrics-api/latest-data/b2696545-7931-47c7-af11-cc2ce1d5e10d
```

#### Response - 201 Success
```json
{
    "status": 201,
    "data": [
        {
            "metric": {
                "__name__": "container_cpu_usage_seconds_total",
                "container_label_com_docker_compose_config_hash": "0564f715245d3d7a8b140ce68be65b5246116cfcc963c1408673c66469716c00",
                "container_label_com_docker_compose_container_number": "1",
                "container_label_com_docker_compose_depends_on": "rate_limiter_service:service_started:false,magnetar:service_started:false,oort:service_started:false,apollo:service_started:false,kuiper:service_started:false,quasar:service_started:false",
                "container_label_com_docker_compose_image": "sha256:bf47588af42892bb70b9ca6a0d41d50928cdd5baa9f984b2bb0937f9a9efc6c4",
                "container_label_com_docker_compose_oneoff": "False",
                "container_label_com_docker_compose_project": "tools",
                "container_label_com_docker_compose_project_config_files": "/home/bunjo/Desktop/college/projects/c12s/tools/docker-compose.yml",
                "container_label_com_docker_compose_project_working_dir": "/home/bunjo/Desktop/college/projects/c12s/tools",
                "container_label_com_docker_compose_service": "lunar-gateway",
                "container_label_com_docker_compose_version": "2.27.0",
                "cpu": "total",
                "id": "/docker/913ef1bcab1e47e25130e74eb499da81e33feaba9bcb38af311a11719553dcbf",
                "image": "lunar-gateway",
                "instance": "health-check:8080",
                "job": "health-check",
                "name": "lunar-gateway",
                "nodeID": "b2696545-7931-47c7-af11-cc2ce1d5e10d"
            },
            "values": [
                [
                    1720514145,
                    "0.034371"
                ]
            ]
        },
        {
            "metric": {
                "__name__": "container_cpu_usage_seconds_total",
                "container_label_com_docker_compose_config_hash": "080eb67bae2cc75484bf1d96130bf808cc64e592f60b5becd19294cacb2e01b5",
                "container_label_com_docker_compose_container_number": "1",
                "container_label_com_docker_compose_image": "sha256:68dd2cdef1732a37435ca3fe75cfa849fbba13f7f0e615a96a1620029a3dad0e",
                "container_label_com_docker_compose_oneoff": "False",
                "container_label_com_docker_compose_project": "tools",
                "container_label_com_docker_compose_project_config_files": "/home/bunjo/Desktop/college/projects/c12s/tools/docker-compose.yml",
                "container_label_com_docker_compose_project_working_dir": "/home/bunjo/Desktop/college/projects/c12s/tools",
                "container_label_com_docker_compose_service": "cassandra",
                "container_label_com_docker_compose_version": "2.27.0",
                "container_label_org_opencontainers_image_ref_name": "ubuntu",
                "container_label_org_opencontainers_image_version": "22.04",
                "cpu": "total",
                "id": "/docker/9efcae5c59f82947c83d60bfb082ccc50c3a0f2ba2830c0798535045ea20a826",
                "image": "tools-cassandra",
                "instance": "health-check:8080",
                "job": "health-check",
                "name": "cassandra",
               

 "nodeID": "b2696545-7931-47c7-af11-cc2ce1d5e10d"
            },
            "values": [
                [
                    1720514145,
                    "21.141999"
                ]
            ]
        },
        {
            "metric": {
                "__name__": "container_cpu_usage_seconds_total",
                "container_label_com_docker_compose_config_hash": "118aa150e527155107a4a55c9d548b48ab1b3a09f2a2d48fe105a93cdf165853",
                "container_label_com_docker_compose_container_number": "1",
                "container_label_com_docker_compose_depends_on": "cassandra:service_healthy:false,nats:service_started:false,vault:service_started:false",
                "container_label_com_docker_compose_image": "sha256:fad6c9e26fe8062a94515924541f15aeef7f35ed2c12f04022b6c1825f53b6b3",
                "container_label_com_docker_compose_oneoff": "False",
                "container_label_com_docker_compose_project": "tools",
                "container_label_com_docker_compose_project_config_files": "/home/bunjo/Desktop/college/projects/c12s/tools/docker-compose.yml",
                "container_label_com_docker_compose_project_working_dir": "/home/bunjo/Desktop/college/projects/c12s/tools",
                "container_label_com_docker_compose_service": "apollo",
                "container_label_com_docker_compose_version": "2.27.0",
                "container_label_restartcount": "2",
                "cpu": "total",
                "id": "/docker/699318abdc21cd8de6c6ca2b18fa08035ede888b96e099de7ff4bc5917029238",
                "image": "apollo",
                "instance": "health-check:8080",
                "job": "health-check",
                "name": "apollo",
                "nodeID": "b2696545-7931-47c7-af11-cc2ce1d5e10d"
            },
            "values": [
                [
                    1720514145,
                    "0.03344"
                ]
            ]
        }
    ]
}
```

##### Response Body Properties (201 Success)
| Property | Type | Description                              |
|----------|------|------------------------------------------|
| status   | int  | HTTP status code                         |
| data     | array| Array of metric objects and their values |

#### Response - 500 Internal Server Error
Server error with error identifier where and when in service error happened.

