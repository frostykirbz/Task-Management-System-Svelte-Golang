<script>
  import axios from 'axios'
  import {onMount} from "svelte"

	let groupname = ""
  let message = ""

  async function CreateGroup(){
    const json = {user_group: groupname}
    console.log(json)
    // try{
      const response = await axios.post("http://localhost:4000/admin-create-group", json)
      console.log(response.data)
    //   if (response)
    //   {
    //     message = response.data.message
    //   }
    // } catch (e){
    //   message = e
    // }
  }

</script>

<style>
	input {
		color: purple;
	}
</style>

<input type="text" bind:value={groupname}>
<p>{message}</p>

<button on:click={CreateGroup}>Calculate</button>
