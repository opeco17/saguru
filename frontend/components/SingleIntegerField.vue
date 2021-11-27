<template>
    <v-text-field
        :value="value"
        @input="$emit('input', $event)"
        outlined
        dense
        single-line
        :label="label"
    ></v-text-field>
</template>

<script>
export default {
    props: {
        value: String,
        label: String
    }
}
</script>
