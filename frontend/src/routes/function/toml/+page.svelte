<script>
	import Sidebar from "$lib/components/Sidebar.svelte";
	import CodeMirror from "svelte-codemirror-editor";

	import Title from '$lib/components/Title.svelte';
    import ErrorBox from '$lib/components/ErrorBox.svelte';
	import Copy from '$lib/components/Copy.svelte';

	let errorText = "";
	let inputText = "";
</script>

<Sidebar />

<Title
	text="TOML"
	description="Validate and format TOML based data"
/>

<ErrorBox text={errorText} />

<div class="flex md:flex md:flex-grow flex-row justify-end space-x-1">
	<Copy text={inputText} />
</div>

<CodeMirror bind:value={inputText} />
