<script lang="ts">
  interface Props {
    placeholder?: string;
  }

  const { placeholder }: Props = $props();
</script>

<input
  type="text"
  class="rounded-md dark:bg-neutral-800 dark:border-neutral-700 text-sm"
  {placeholder}
/>
