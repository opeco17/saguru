<template>
    <v-btn
        class="mr-4 white--text"
        color="#F85758"
        depressed
        :disabled="disabled"
        @click="$emit('click', $event)"
    >
        <slot></slot>
    </v-btn>
</template>

<script>
export default {
    props: {
        disabled: Boolean,
    }
}
</script>

