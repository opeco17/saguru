<template>
  <v-select
    :value="value"
    @input="$emit('input', $event)"
    :items="items"
    outlined
    single-line
    dense
    :menu-props="{offsetY: true, maxHeight: 250, bottom: true}"
  ></v-select>
</template>

<script>
export default {
  props: {
    value: String,
    items: Array
  }
}
</script>
