<script>
  import axios from "axios";
  import { Button, Modal, ModalBody, ModalHeader, ModalFooter, Table, Row } from "sveltestrap";
  import Navbar from "../Navbar/IsLoggedInAdmin.svelte";
  import AddGroup from "./AddGroup.svelte";
  import AddUserToGroup, { handleAddUserToGroup } from "./AddUserToGroup.svelte";

  let usersGroupData = [];
  let groupname = "";
  let openModalAddGroup = false;
  let openModalAddUserToGroup = false;
  const size = "lg";

  $: getGroupInfo();

  async function getGroupInfo() {
    try {
      const response = await axios.get("http://localhost:4000/get-users-in-group");
      if (response) {
        console.log(response.data)
        usersGroupData = response.data;
      }
    } catch (error) {
      console.log(error);
    }
  }
  
  const toggleAddGroup = (e) => {
    e.preventDefault();
    openModalAddGroup = !openModalAddGroup;
    getGroupInfo();
    groupname = "";
  };

  const toggleAddUserToGroup = (e) => {
    e.preventDefault();
    openModalAddUserToGroup = !openModalAddUserToGroup;
    getGroupInfo();
  };
</script>

<style>
</style>

<Navbar />

<div>
  <h3>Group Management</h3>
  <Button color="primary" on:click={toggleAddGroup}>Add Group</Button>
  <Button color="primary" on:click={toggleAddUserToGroup}>Add User To Group</Button>

  <br/><br/>

  <Table bordered>
    <thead>
      <tr>
        <th>Groupname</th>
        <th># of people</th>
      </tr>
    </thead>
    <tbody>
      {#each usersGroupData as userGroupData}
      <tr>
        <td>{userGroupData.user_group}</td>
        <td>{userGroupData.user_count}</td>
      </tr>
      {/each}
    </tbody>
  </Table>
 
  <!-- Modal: Add Group  -->
  <Modal isOpen={openModalAddGroup} {toggleAddGroup} {size}>
    <ModalHeader {toggleAddGroup}>Add Group</ModalHeader>
    <ModalBody>
      <AddGroup {groupname}/>
    </ModalBody>

    <ModalFooter>
      <Button on:click={handleAddGroup} style="background-color: #FCA311; border: none;">Add</Button>
      <Button class="back-button" color="danger" on:click={toggleAddGroup}>Back</Button>
    </ModalFooter>
  </Modal>
  
  <!-- Modal: Add User To Group  -->
  <Modal isOpen={openModalAddUserToGroup} {toggleAddUserToGroup} {size}>
    <ModalHeader {toggleAddUserToGroup}>Add User To Group</ModalHeader>
    <ModalBody>
      <AddUserToGroup />
    </ModalBody>

    <ModalFooter>
      <Button on:click={handleAddUserToGroup} style="background-color: #FCA311; border: none;">Add</Button>
      <Button class="back-button" color="danger" on:click={toggleAddUserToGroup}>Back</Button>
    </ModalFooter>
  </Modal>
</div>
