package main

import (
  "context"
  "github.com/hashicorp/terraform-plugin-framework/tfsdk"
  "terraform-provider-chirpstack/chirpstack"
)

func main() {
  tfsdk.Serve(context.Background(), chirpstack.New, tfsdk.ServeOpts{
    Name: "chirpstack",
  })
}
