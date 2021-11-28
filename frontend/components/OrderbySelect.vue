<template>
  <v-select
    :value="value"
    @input="$emit('input', $event)"
    :items="items"
    outlined
    single-line
    dense
  >
    <template v-slot:selection="{ item }">
      {{ item | formatOrderMetricLabel }}<v-icon>{{ item | orderMetricIcon }}</v-icon>
    </template>
    <template v-slot:item="{ item }">
      {{ item | formatOrderMetricLabel }}<v-icon>{{ item | orderMetricIcon }}</v-icon>
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
      value = value.replace('_asc', '').replace('_desc', '').replace("_", " ")
      value = value.charAt(0).toUpperCase() + value.slice(1)
      return value
    },
    orderMetricIcon(value) {
      if (value.includes('desc')) {
        return "mdi-menu-up"
      } else {
        return "mdi-menu-down"
      }
    }
  }
}
</script>
