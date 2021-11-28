<template>
  <v-autocomplete
    :value="value"
    @input="$emit('input', $event)"
    :items="items"
    outlined
    chips
    small-chips
    single-line
    multiple
  >
    <template v-slot:selection="data">
      <v-chip
        v-bind="data.attrs"
        :input-value="data.selected"
        close
        outlined
        color="grey darken-2"
        small
        @click="data.select"
        @click:close="$emit('close', data.item)"
      >
        {{ data.item }}
      </v-chip>
    </template>
  </v-autocomplete>
</template>

<script>
export default {
  props: {
    value: Array,
    items: Array
  }
}
</script>
