<template>
  <div>
    <v-autocomplete
      :value="value"
      @input="$emit('input', $event)"
      :items="items"
      outlined
      chips
      small-chips
      single-line
      multiple
      :menu-props="{offsetY: true, bottom: true}"
      class="hidden-xs-only"
    >
      <template v-slot:selection="data">
        <multiple-chip :data="data" @close="$emit('close', data.item)" />
      </template>
    </v-autocomplete>

    <v-select
      :value="value"
      @input="$emit('input', $event)"
      :items="items"
      outlined
      chips
      small-chips
      single-line
      multiple
      :menu-props="{offsetY: true, maxHeight: 250, bottom: true}"
      class="hidden-sm-and-up"
    >
      <template v-slot:selection="data">
        <multiple-chip :data="data" @close="$emit('close', data.item)" />
      </template>
    </v-select>
  </div>
</template>

<script>
export default {
  props: {
    value: Array,
    items: Array
  }
}
</script>
