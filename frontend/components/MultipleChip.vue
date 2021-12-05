<template>
  <v-chip
    v-bind="data.attrs"
    close
    outlined
    color="grey darken-2"
    small
    @click:close="$emit('close')"
    >
      {{ data.item }}
  </v-chip>
</template>

<script>
export default {
  props: {
    data: Object
  }
}
</script>
