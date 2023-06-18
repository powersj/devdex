<script>
    import Header from "./sidebar/Header.svelte";
    import Item from "./sidebar/Item.svelte";

    import IconBinary from "@tabler/icons-svelte/dist/svelte/icons/IconBinary.svelte";
    import IconBracketsContain from "@tabler/icons-svelte/dist/svelte/icons/IconBracketsContain.svelte";
    import IconCode from "@tabler/icons-svelte/dist/svelte/icons/IconCode.svelte";
    import IconCodeAsterix from "@tabler/icons-svelte/dist/svelte/icons/IconCodeAsterix.svelte";
    import IconCsv from "@tabler/icons-svelte/dist/svelte/icons/IconCsv.svelte";
    import IconFileSettings from "@tabler/icons-svelte/dist/svelte/icons/IconFileSettings.svelte";
    import IconHash from "@tabler/icons-svelte/dist/svelte/icons/IconHash.svelte";
    import IconHome from "@tabler/icons-svelte/dist/svelte/icons/IconHome.svelte";
    import IconCodeDots from "@tabler/icons-svelte/dist/svelte/icons/IconCodeDots.svelte";
    import IconDatabaseSearch from "@tabler/icons-svelte/dist/svelte/icons/IconDatabaseSearch.svelte";
    import IconTypography from "@tabler/icons-svelte/dist/svelte/icons/IconTypography.svelte";

    import IconSourceCode from "@tabler/icons-svelte/dist/svelte/icons/IconSourceCode.svelte";
    import IconForms from "@tabler/icons-svelte/dist/svelte/icons/IconForms.svelte";
    import IconFileDiff from "@tabler/icons-svelte/dist/svelte/icons/IconFileDiff.svelte";
    import IconRegex from "@tabler/icons-svelte/dist/svelte/icons/IconRegex.svelte";
    import IconCertificate from "@tabler/icons-svelte/dist/svelte/icons/IconCertificate.svelte";

    import IconKey from "@tabler/icons-svelte/dist/svelte/icons/IconKey.svelte";
    import IconClock from "@tabler/icons-svelte/dist/svelte/icons/IconClock.svelte";
    import IconRulerMeasure from "@tabler/icons-svelte/dist/svelte/icons/IconRulerMeasure.svelte";
    import IconNetwork from "@tabler/icons-svelte/dist/svelte/icons/IconNetwork.svelte";
    import IconPalette from "@tabler/icons-svelte/dist/svelte/icons/IconPalette.svelte";

    import IconCpu2 from "@tabler/icons-svelte/dist/svelte/icons/IconCpu2.svelte";
    import IconQrcode from "@tabler/icons-svelte/dist/svelte/icons/IconQrcode.svelte";
    import IconAdjustmentsSearch from "@tabler/icons-svelte/dist/svelte/icons/IconAdjustmentsSearch.svelte";
    import IconBinaryTree from "@tabler/icons-svelte/dist/svelte/icons/IconBinaryTree.svelte";
</script>

<aside id="default-sidebar" class="fixed top-0 left-0 z-40 w-64 h-screen transition-transform -translate-x-full sm:translate-x-0" aria-label="Sidebar">
  <div class="min-h-screen flex flex-col flex-auto flex-shrink-0 antialiased bg-gray-50 text-gray-800">
    <div class="fixed flex flex-col top-0 left-0 w-64 bg-white h-full border-r">
      <div class="overflow-y-auto overflow-x-hidden flex-grow">
        <ul class="flex flex-col py-4 space-y-1">
          <li class="px-5">
            <div class="flex flex-row items-center h-8">
              <div class="font-light tracking-wide text-gray-500">devdex</div>
            </div>
          </li>

          <Item text="Home" icon={IconHome} href="/" />

          <Header text="Configuration" />
          <Item text="JSON" icon={IconCodeDots} href="/function/json" />
          <Item text="JSON Query" icon={IconCodeAsterix} href="/function/json-query" />
          <Item text="TOML" icon={IconBracketsContain} href="/function/toml" />
          <Item text="YAML" icon={IconFileSettings} href="/function/yaml" />
          <Item text="XML" icon={IconCode} href="/function/xml" />
          <Item text="CSV" icon={IconCsv} href="/function/csv" />
          <Item text="SQL" icon={IconDatabaseSearch} href="/function/sql" />

          <Header text="Strings" />
          <Item text="Modify" icon={IconTypography} href="/function/string" />
          <Item text="Hash Generator" icon={IconHash} href="/function/hash" />
          <Item text="HTML Encode/Decode" icon={IconSourceCode} href="/function/html" />
          <Item text="URL Encode/Decode" icon={IconForms} href="/function/url" />
          <Item text="Diff" icon={IconFileDiff} href="/function/diff" />
          <Item text="Regex" icon={IconRegex} href="/function/regex" />

          <Header text="Tools" />
          <Item text="Base Conversion" icon={IconBinary} href="/function/base" />
          <Item text="Certificate" icon={IconCertificate} href="/function/certificates" />
          <Item text="Passwords" icon={IconKey} href="/function/passwords" />
          <Item text="Crontab" icon={IconClock} href="/function/crontab" />
          <Item text="Measurements" icon={IconRulerMeasure} href="/function/measurements" />
          <Item text="Data" icon={IconBinary} href="/function/data" />
          <Item text="IP Address" icon={IconNetwork} href="/function/ip-address" />
          <Item text="Color" icon={IconPalette} href="/function/color" />
          <Item text="Timestamp" icon={IconClock} href="/function/timestamp" />
          <Item text="UUID" icon={IconCpu2} href="/function/uuid" />
          <Item text="QR Code" icon={IconQrcode} href="/function/qrcode" />

          <Header text="Reference" />
          <Item text="Algorithms" icon={IconAdjustmentsSearch} href="/function/algorithms" />
          <Item text="Data Structures" icon={IconBinaryTree} href="/function/data_structures" />

        </ul>
      </div>
    </div>
  </div>
</aside>
