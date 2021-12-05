<template>
  <v-select
    :value="value"
    @input="$emit('input', $event)"
    :items="items"
    outlined
    single-line
    dense
    :menu-props="{offsetY: true, maxHeight: 250, bottom: true}"
  >
    <template v-slot:selection="{ item }" dense>
      {{ item | formatOrderMetricLabel }}
    </template>
    <template v-slot:item="{ item }">
      {{ item | formatOrderMetricLabel }}
    </template>
  </v-select>
</template>

<script>
export default {
  props: {
    value: String,
    items: Array
  },
  filters: {
    formatOrderMetricLabel(value) {
      let prefix = value.includes('desc') ? "Most" : "Fewest"
      value = value.replace('_asc', '').replace('_desc', '').replace("_count", "s")
      return prefix + " " + value
    },
  }
}
</script>
