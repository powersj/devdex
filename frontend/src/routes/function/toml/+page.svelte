<script lang="ts">
	import CodeMirror from "svelte-codemirror-editor";

	import { FormatTOML } from "$lib/wailsjs/go/main/App.js";

	import Title from '$lib/components/Title.svelte';
	import Button from '$lib/components/Button.svelte';
	import ErrorBox from '$lib/components/ErrorBox.svelte';
	import Copy from '$lib/components/Copy.svelte';

	let errorText = "";
	let inputText = "";

	function format() {
		FormatTOML(inputText).then((result: string) => {
			inputText = result
			errorText = ""
		}).catch(function (e: any) {
			errorText = e
		})
    }
</script>

<Title
	text="TOML"
	description="Validate and format TOML based data"
/>

<ErrorBox text={errorText} />

<div class="flex md:flex md:flex-grow flex-row justify-end space-x-1">
	<Button text="format" on:click={format} />
	<Copy text={inputText} />
</div>

<CodeMirror
	bind:value={inputText}
	styles={{
		"&": {height: "100vh"},
		".cm-scroller": {overflow: "auto"}
	}}
/>
