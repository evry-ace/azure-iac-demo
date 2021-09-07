resource "azurerm_frontdoor" "lb" {
  name                = data.azurerm_resource_group.main.name
  resource_group_name = data.azurerm_resource_group.main.name

  enforce_backend_pools_certificate_name_check = false

  routing_rule {
    name               = "app"
    accepted_protocols = ["Http", "Https"]
    patterns_to_match  = ["/*"]
    frontend_endpoints = ["app-frontend"]
    forwarding_configuration {
      forwarding_protocol = "HttpOnly"
      backend_pool_name   = "app-backend"
    }
  }

  backend_pool_load_balancing {
    name = "app-settings"
  }

  backend_pool_health_probe {
    name = "app-health"
  }

  backend_pool {
    name = "app-backend"
    backend {
      host_header = azurerm_container_group.app.fqdn
      address     = azurerm_container_group.app.ip_address
      http_port   = 80
      https_port  = 443
    }

    load_balancing_name = "app-settings"
    health_probe_name   = "app-health"
  }

  frontend_endpoint {
    name      = "app-frontend"
    host_name = "${data.azurerm_resource_group.main.name}.azurefd.net"
  }
}

output "frontdoor_dns" {
  value = azurerm_frontdoor.lb.frontend_endpoint.0.host_name
}
