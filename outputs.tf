output "cluster_name" {
  value = aws_ecs_cluster.blockchain_cluster.name
}

output "service_name" {
  value = aws_ecs_service.blockchain_service.name
}