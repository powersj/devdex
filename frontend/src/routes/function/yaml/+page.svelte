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
	text="YAML"
	description="Validate and format YAML based data"
/>

<ErrorBox text={errorText} />

<Copy text={inputText} />
<CodeMirror bind:value={inputText} />
