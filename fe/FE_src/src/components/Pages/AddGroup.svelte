<script>
    import axios from "axios";
    import { errorToast, successToast } from "../toast";
    import { Form, FormGroup, Input, Label, Button, Modal, ModalHeader, ModalFooter, Col, Row, Spinner, ModalBody, Styles } from "sveltestrap";
    
    export let groupname;

    export async function handleAddGroup(e) {
      e.preventDefault();
      const json = { user_group: groupname };
      console.log(json);
      try {
        const response = await axios.post("http://localhost:4000/admin-create-group", json, { withCredentials: true });
            
        // code 200: error message
        // code 201: success message
        if (response.data.code == 200) {
          errorToast(response.data.message);
          groupValue = "";
        }
        else if (response.data.code == 201) {
          successToast(response.data.message);
          groupValue = "";
        }
      } catch (e) {
        console.log(e);
      }
    } 
</script>
  
<style>
    /* input {
      color: purple;
    } */
</style>

<Form on:submit={handleAddGroup}>
  <FormGroup>
    <label for="group">Groupname</label>
    <Input placeholder="groupname" type="text" bind:value={groupname} autofocus />
  </FormGroup>
</Form>