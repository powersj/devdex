<script>
	import Sidebar from "$lib/components/Sidebar.svelte";
    import "../app.css";
</script>

<Sidebar />

<main>
	<div class="p-4 sm:ml-64">
		<div class="p-4 border-2 rounded-lg border-gray-200  dark:border-gray-700">
            <slot />
        </div>
    </div>
</main>
