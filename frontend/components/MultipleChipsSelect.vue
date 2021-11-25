<template>
    <v-autocomplete
        v-model="model"
        :items="items"
        outlined
        chips
        small-chips
        single-line
        multiple
        clearable
    ></v-autocomplete>
</template>

<script>
export default {
    props: {
        model: Array,
        items: Array
    }
}
</script>
