<script>
	//export let name;
	import {Router, Route} from "svelte-routing"
	import ProtectedRoute from "./ProtectedRoute.svelte"
	import Login from "./components/Login.svelte"
	import Home from "./components/Home.svelte"
	import MgtUser from "./components/MgtUser.svelte"
</script>

<main>
	<Router>
		<Route path="/" component={Login} />
		<Route path="/home" component={Home}/>
		<ProtectedRoute path="/user" component={MgtUser} />
	</Router>
</main>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}
/* 
	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	} */

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>