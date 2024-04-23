output "module_alb_dns_name" {
  value       = module.alb.lb_dns_name
  description = "The DNS name of the Application Load Balancer created by the module"
}
