<!--
page_title: Azure Terraform Suffering
page_description: "A painful problem I've found in Azure, with an interesting solution..."
page_date: 2026-07-18
page_image: https://julianorchard.co.uk/res/default.png
rel: posts/{date}/{name}
draft: true
-->

# Azure Terraform Pain

Azure is a cloud provider that you *can use*. But it has some funny quirks!
Here's a fun (and interesting) issue I've encountered and been meaning to write
about for a while...

## The Problem

> For a quick TL;DR of the problem, read
> [this](https://github.com/hashicorp/terraform-provider-azurerm/issues/6098#issuecomment-1156049806)
> GitHub issue

There's a good many reasons to use [Azure VM
Extensions](https://learn.microsoft.com/en-us/azure/virtual-machines/extensions/overview).
In *our* case, the reason is that we wanted to automate the addition of a VM as
an [AVD Session
Host](https://learn.microsoft.com/en-us/azure/virtual-desktop/add-session-hosts-host-pool?tabs=portal%2Cgui&pivots=host-pool-standard).
This is a pretty well known pattern that involves running some
[PowerShell
scripts](https://github.com/Azure/RDS-Templates/blob/10a904c0cf2c566941251d24b18d505769b11d36/wvd-sh/terraform-azurerm-azuresvirtualdesktop/archive/host.tf#L95C1-L126C2)
and doing a VM restart.

However: **when the VM is off, the Azure API [cannot interact with th
resource
anymore](https://github.com/hashicorp/terraform-provider-azurerm/issues/6098#issuecomment-1156049806)**.
The [AVD Scaling
Plans](https://learn.microsoft.com/en-us/azure/virtual-desktop/autoscale-scenarios)
(fairly misleading in name: they just power on and off the VMs) cause
Terraform to fail in destroying the VM Extension resources.

## The Solution

My workaround was simply to use a `null_resource` with `when = destroy`:

```hcl
resource "null_resource" "this" {
  triggers = {
    resource_group_name = var.resource_group_name
    computer_name        = azurerm_windows_virtual_machine.this.computer_name
  }

  provisioner "local-exec" {
    when    = destroy
    command = "az vm start -g ${self.triggers.resource_group_name} -n ${self.triggers.computer_name}"
    # We exclusively run this through Windows with PowerShell
    interpreter = ["powershell", "-Command"]
  }

  depends_on = [
    azurerm_virtual_machine_extension.restart
  ]
}
```

Before we delete the VM, we turn it back on using the Az CLI `az vm start`
command! This works pretty well, and for quite a large number of Session Hosts.

We've never experienced any issues which required this, but adding a check to
make sure the VM actually started would probably be a nice quality of life
addition. It was just unnecessary in our case!

## Notes

I came up with this solution in 2023: however, [this solution by Jose
Espitia](https://github.com/hashicorp/terraform-provider-azurerm/issues/17262#issuecomment-1924565954)
(which I didn't see at the time) looks like a great approach too!
