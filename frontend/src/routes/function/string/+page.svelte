<script lang="ts">
    import CodeMirror from "svelte-codemirror-editor";

    import Button from '$lib/components/Button.svelte';
	import Title from '$lib/components/Title.svelte';
    import ErrorBox from '$lib/components/ErrorBox.svelte';
	import Copy from '$lib/components/Copy.svelte';

    let errorText = "";
    let inputText = "";

    function unique() {}
    function sort() {}
    function reverseSort() {}
    function toLower() {}
    function toUpper() {}
    function shuffle() {}
    function trailingWhitespace() {}
    function leadingWhitespace() {}
    function newlines() {}
    function replaceAll() {}
</script>

<Title
	text="String"
	description="Modify strings directly"
/>

<ErrorBox text={errorText} />

<h3>modifiers</h3>
<Button text="unique" on:click={unique} />
<Button text="sort" on:click={sort} />
<Button text="revere sort" on:click={reverseSort} />
<Button text="to lower" on:click={toLower} />
<Button text="to upper" on:click={toUpper} />
<Button text="shuffle" on:click={shuffle} />

<h3>whitespace</h3>
<Button text="trailing whitespace" on:click={trailingWhitespace} />
<Button text="leading whitespace" on:click={leadingWhitespace} />
<Button text="newlines" on:click={newlines} />

<h3>find/replace</h3>
<input name="find"/>
<input name="replace"/>
<Button text="replace all" on:click={replaceAll} />

<Copy text={inputText} />
<CodeMirror bind:value={inputText} />
