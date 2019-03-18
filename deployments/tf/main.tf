provider "google" {
  project = "gbookshelf-dev"
  region  = "asia-northeast1"
  zone    = "a"
}

# https://www.terraform.io/docs/providers/kubernetes/index.html
provider "kubernetes" {
  # host = "https://104.196.242.174"

  client_certificate     = "${base64decode(google_container_cluster.gbookshelf.master_auth.0.client_certificate)}"
  client_key             = "${base64decode(google_container_cluster.gbookshelf.master_auth.0.client_key)}"
  cluster_ca_certificate = "${base64decode(google_container_cluster.gbookshelf.master_auth.0.cluster_ca_certificate)}"
  load_config_file       = "false"
}

locals {
  project_name    = "gbookshelf"
  env             = "dev"
  resource_prefix = "${local.project_name}-${local.env}"
  region          = "asia-northeast1"
  machine_type    = "n1-standard-1"
}

data "google_container_engine_versions" "asia_northeast1a" {
  zone = "asia-northeast1-a"
}

resource "google_container_cluster" "gbookshelf" {
  name   = "${local.resource_prefix}"
  region = "${local.region}"

  # We can't create a cluster with no node pool defined, but we want to only use
  # separately managed node pools. So we create the smallest possible default
  # node pool and immediately delete it.
  remove_default_node_pool = true

  initial_node_count = 1

  # min_master_version = "${data.google_container_engine_versions.asia_northeast1a.latest_node_version}"
  # node_version       = "${data.google_container_engine_versions.asia_northeast1a.latest_node_version}"

  min_master_version = "1.12.5-gke.5"
  node_version       = "1.12.5-gke.5"
  # Setting an empty username and password explicitly disables basic auth
  master_auth {
    username = ""
    password = ""
  }
  node_config {
    oauth_scopes = [
      "https://www.googleapis.com/auth/compute",
      "https://www.googleapis.com/auth/devstorage.read_only",
      "https://www.googleapis.com/auth/logging.write",
      "https://www.googleapis.com/auth/monitoring",
    ]

    labels = {
      project = "${local.project_name}"
      env     = "${local.env}"
    }

    tags = ["${local.project_name}", "${local.env}"]
  }
}

resource "google_container_node_pool" "gbookshelf_preemptible_nodes" {
  name       = "${local.resource_prefix}-pool"
  region     = "${local.region}"
  cluster    = "${google_container_cluster.gbookshelf.name}"
  node_count = 1

  node_config {
    preemptible  = true
    machine_type = "${local.machine_type}"

    oauth_scopes = [
      "https://www.googleapis.com/auth/compute",
      "https://www.googleapis.com/auth/devstorage.read_only",
      "https://www.googleapis.com/auth/logging.write",
      "https://www.googleapis.com/auth/monitoring",
    ]
  }
}

# The following outputs allow authentication and connectivity to the GKE Cluster
# by using certificate-based authentication.
output "client_certificate" {
  value = "${google_container_cluster.gbookshelf.master_auth.0.client_certificate}"
}

output "client_key" {
  value = "${google_container_cluster.gbookshelf.master_auth.0.client_key}"
}

output "cluster_ca_certificate" {
  value = "${google_container_cluster.gbookshelf.master_auth.0.cluster_ca_certificate}"
}
