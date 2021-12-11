<template>
  <v-btn
    class="mr-4"
    color="#F85758"
    outlined
    :width="width"
    :disabled="disabled"
    @click="$emit('click', $event)"
  >
    <slot></slot>
  </v-btn>
</template>

<script>
export default {
  props: {
    width: String,
    disabled: Boolean,
  }
}
</script>

