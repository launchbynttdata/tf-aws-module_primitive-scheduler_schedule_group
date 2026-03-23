logical_product_family  = "launch"
logical_product_service = "scheduler"
class_env               = "dev"
instance_env            = 0
instance_resource       = 1

resource_names_map = {
  schedule_group = { name = "schedulergroup1", max_length = 64 }
  scheduler_role = { name = "iamrole1", max_length = 64 }
  schedule_1     = { name = "sched1", max_length = 64 }
  schedule_2     = { name = "sched2", max_length = 64 }
}

tags = {
  Environment = "test"
  Terraform   = "true"
}
