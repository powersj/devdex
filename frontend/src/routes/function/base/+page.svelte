<h1 class="text-2xl">Base Converter</h1>
<p>Used to convert a value from one base to another</p>

<div class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative" role="alert">
    {errorText}
</div>

<p>Binary (base 2)</p>
<input bind:value="{base2}" on:input={calcBase2} type="text" />

<p>Octal (base 8)</p>
<input bind:value="{base8}" on:input={calcBase8} type="text" />

<p>Decimal (base 10)</p>
<input bind:value="{base10}" on:input={calcBase10} type="text" />

<p>Hex (base 16)</p>
<input bind:value="{base16}" on:input={calcBase16} type="text" />

<p>
    Other base:
    <select bind:value="{baseOther}" on:select={calcBaseOther}>
        <option></option>
        <option>3</option>
        <option>4</option>
        <option>5</option>
        <option>6</option>
        <option>7</option>
        <option>9</option>
    </select>
</p>
<input on:input={calcBaseOther} type="text" />

<script lang="ts">
    import { Base2, Base8, Base10, Base16 } from "$lib/wailsjs/go/main/App.js";

    let base2 = ""
    let base8 = ""
    let base10 = ""
    let base16 = ""
    let baseOther = ""
    let errorText = ""

    function calcBase2() {
        Base8(base2).then((result: string) => {
            base8 = result
        }).catch((e: any) => errorText = e);
        Base10(base2).then((result: string) => {
            base10 = result
        }).catch((e: any) => errorText = e);
        Base16(base2).then((result: string) => {
            base16 = result
        }).catch((e: any) => errorText = e);
    }

    function calcBase8() {
        Base2(base8).then((result: string) => {
            base2 = result
        }).catch((e: any) => errorText = e);
        Base10(base8).then((result: string) => {
            base10 = result
        }).catch((e: any) => errorText = e);
        Base16(base8).then((result: string) => {
            base16 = result
        }).catch((e: any) => errorText = e);
    }

    function calcBase10() {
        Base2(base10).then((result: string) => {
            base2 = result
        }).catch((e: any) => errorText = e);
        Base8(base10).then((result: string) => {
            base8 = result
        }).catch((e: any) => errorText = e);
        Base16(base10).then((result: string) => {
            base16 = result
        }).catch((e: any) => errorText = e);
    }

    function calcBase16() {
        Base2(base16).then((result: string) => {
            base2 = result
        }).catch((e: any) => errorText = e);
        Base8(base16).then((result: string) => {
            base8 = result
        }).catch((e: any) => errorText = e);
        Base10(base16).then((result: string) => {
            base10 = result
        }).catch((e: any) => errorText = e);
    }

    function calcBaseOther() {
        Base2(baseOther).then((result: string) => {
            base2 = result
        }).catch((e: any) => errorText = e);
        Base8(baseOther).then((result: string) => {
            base8 = result
        }).catch((e: any) => errorText = e);
        Base10(baseOther).then((result: string) => {
            base10 = result
        }).catch((e: any) => errorText = e);
        Base16(baseOther).then((result: string) => {
            base10 = result
        }).catch((e: any) => errorText = e);
    }

</script>
