<script>
	import Sidebar from "$lib/components/Sidebar.svelte";
    import { page } from '$app/stores';
</script>

<Sidebar />

<main>
    <h1>{$page.status}: {$page.error.message}</h1>
</main>
