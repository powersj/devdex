<!--
    https://www.jstoolset.com/json-query
    https://www.influxdata.com/versions.json
-->

<script lang="ts">
    import CodeMirror from "svelte-codemirror-editor";
    import { json } from "@codemirror/lang-json";

    import Title from '$lib/components/Title.svelte';
	import Button from '$lib/components/Button.svelte';
    import ErrorBox from '$lib/components/ErrorBox.svelte';
	import Copy from '$lib/components/Copy.svelte';

    let inputText = "";
    let queryText = "";
    let errorText = "";
    let resultText = "";

    function minify(){

    }

    function beautify(){

    }
</script>

<Title
	text="JSON Query"
	description="Query a JSON string"
/>

<ErrorBox text={errorText} />

<div class="grid grid-cols-3 gap-3">
    <div>
        Input
        <div class="flex md:flex md:flex-grow flex-row justify-end space-x-1">
            <Button text="minify" on:click={minify} />
            <Button text="beautify" on:click={beautify} />
        </div>
    </div>
    <div>Query</div>
    <div>
        Output
        <div class="flex md:flex md:flex-grow flex-row justify-end space-x-1">
            <Button text="minify" on:click={minify} />
            <Button text="beautify" on:click={beautify} />
            <Copy text={resultText} />
        </div>
    </div>
</div>

<div class="grid grid-cols-3 gap-3">
    <CodeMirror bind:value={inputText} lang={json()}
    styles={{
		"&": {height: "100vh"},
		".cm-scroller": {overflow: "auto"}
	}}
    />
    <CodeMirror bind:value={queryText} lang={json()}
    styles={{
		"&": {height: "100vh"},
		".cm-scroller": {overflow: "auto"}
	}}/>
    <CodeMirror bind:value={resultText} lang={json()}
    styles={{
		"&": {height: "100vh"},
		".cm-scroller": {overflow: "auto"}
	}} />
</div>
