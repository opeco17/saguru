<template>
    <v-select
        :value="value"
        @input="$emit('input', $event)"
        :items="items"
        outlined
        single-line
        dense
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
