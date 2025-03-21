<script lang="ts">
  import { Meta } from "$lib/components";
  import InputField from "$lib/components/InputField.svelte";
  import type { PageData } from "./$types";

  let { data }: { data: PageData } = $props();
</script>

<Meta title="Search page" />
<div class="min-h-[90dvh] flex justify-center items-start">
  <div>
    <InputField />
  </div>
</div>
