<script lang="ts">
	import CodeMirror from "svelte-codemirror-editor";
	import { json } from "@codemirror/lang-json";

	import { FormatJSON, MinifyJSON } from "$lib/wailsjs/go/main/App.js";

	import Title from '$lib/components/Title.svelte';
	import Button from '$lib/components/Button.svelte';
	import ErrorBox from '$lib/components/ErrorBox.svelte';
	import Copy from '$lib/components/Copy.svelte';

    let errorText = "";
    let inputText = "";

    function format() {
		try {
			JSON.parse(inputText)
		} catch(e: any) {
			errorText = 'invalid JSON: ' + e.message
		}

		FormatJSON(inputText).then((result: string) => {
			inputText = result
			errorText = ""
		}).catch((e: any) => errorText = e);
    }
    function minify() {
		try {
			JSON.parse(inputText)
		} catch(e: any) {
			errorText = 'invalid JSON: ' + e.message
		}

        MinifyJSON(inputText).then((result: string) => {
            inputText = result
			errorText = ""
        }).catch((e: any) => errorText = e);
    }
</script>


<Title
	text="JSON"
	description="Use this function to validate JSON as well as minify or beautify JSON"
/>

<ErrorBox text={errorText} />

<div class="flex md:flex md:flex-grow flex-row justify-end space-x-1">
	<Button text="format" on:click={format} />
	<Button text="minify" on:click={minify} />
	<Copy text={inputText} />
</div>

<CodeMirror
	bind:value={inputText}
	lang={json()}
	styles={{
		"&": {height: "100vh"},
		".cm-scroller": {overflow: "auto"}
	}}
/>
