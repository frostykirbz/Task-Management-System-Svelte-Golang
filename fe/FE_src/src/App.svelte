<script>
	// export let name;
	import { SvelteToast } from '@zerodevx/svelte-toast'
	import { Router } from "@roxi/routify";
  import { routes } from "@roxi/routify/tmp/routes";
</script>

<!-- <main>
	<SvelteToast />
	<h1>Hello {name}!</h1>
	<p>Visit the <a href="https://svelte.dev/tutorial">Svelte tutorial</a> to learn how to build Svelte apps.</p>
</main> -->

<main>
	<Router {routes}/>
</main>
<!-- <main>
	<SvelteToast/>
	<Router>
		<nav>
			<Link to="home">Home</Link>
		</nav>
		<route path={"home"} component={Home}/>
	</Router>
</main> -->

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>