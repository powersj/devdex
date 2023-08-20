<script lang="ts">
	import { Beautify, Minify } from "$lib/wailsjs/go/main/App.js";
	import { json } from "@codemirror/lang-json";
	import CodeMirror from "svelte-codemirror-editor";

	import Title from '$lib/components/Title.svelte';
	import Button from '$lib/components/Button.svelte';
    import ErrorBox from '$lib/components/ErrorBox.svelte';
	import Copy from '$lib/components/Copy.svelte';

    let errorText = "";
    let inputText = "";

    function beautify() {
		try {
			JSON.parse(inputText)
		} catch(e: any) {
			errorText = 'Error parsing: ' + e.message
		}

		Beautify(inputText).then((result: string) => {
			inputText = result
			errorText = ""
		}).catch((e: any) => errorText = 'Error parsing: ' + e.message);
    }
    function minify() {
		try {
			JSON.parse(inputText)
		} catch(e: any) {
			errorText = 'Error parsing: ' + e.message
		}

        Minify(inputText).then((result: string) => {
            inputText = result
			errorText = ""
        }).catch((e: any) => errorText = e);
    }
    function treeView() {
        errorText = "tree view is not implemented";
    }
	function stringify() {
        errorText = "stringify is not implemented";
    }
</script>


<Title
	text="JSON"
	description="Use this function to validate JSON as well as minify or beautify JSON"
/>

<ErrorBox text={errorText} />

<div class="flex md:flex md:flex-grow flex-row justify-end space-x-1">
	<Button text="beautify" on:click={beautify} />
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
