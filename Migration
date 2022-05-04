resource "random_string" "identifier-suffix" {
  length = "3"
  special = "false"
  upper = "false"
}

resource "aws_elasticache_subnet_group" "redis" {
  name = "codecov-redis"
  subnet_ids = ["${var.vpc_private_subnet_ids}"]
}

resource "aws_security_group" "elasticache" {
  name_prefix = "codecov-elasticache"
  vpc_id      = "${var.vpc_id}"

  ingress {
    from_port = 6379
    to_port   = 6379
    protocol  = "tcp"

    cidr_blocks = [
      "${data.aws_vpc.vpc.cidr_block}",
    ]
  }
}

resource "aws_elasticache_cluster" "redis" {
  cluster_id = "codecov-redis-${random_string.identifier-suffix.result}"
  engine = "redis"
  node_type = "${var.redis_node_type}"
  num_cache_nodes = "${var.redis_num_nodes}"
  engine_version = "5.0.3"
  subnet_group_name = "${aws_elasticache_subnet_group.redis.name}"
  security_group_ids = ["${aws_security_group.elasticache.id}"]
}
