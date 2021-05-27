func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
        return resourceServerRead(d, m)
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
        return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
        return resourceServerRead(d, m)
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
        return nil
}
