import { writable } from "svelte/store"

export const arrayWritable = <Data = unknown>() => {
  const { subscribe, update, set } = writable<Data[]>([])

  return {
    subscribe,
    update,
    deleteFromIndex: (...indices: number[]) => update((prevData) => {
      return prevData.filter((_, index) => !indices.includes(index))
    }),
    clear: () => set([])
  }
}