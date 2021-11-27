<template>
    <v-text-field
        :value="value"
        @input="$emit('input', $event)"
        outlined
        dense
        single-line
        :rules="customRules"
        :label="label"
    ></v-text-field>
</template>

<script>
export default {
    props: {
        value: String,
        label: String,
        rules: Array,
    },
    computed: {
        customRules() {
            let rules = [
                v => !(v !== '' && isNaN(v)) || "Must be number",
                v => !(v !== '' && Number(v) < 0) || "Must be greater than or equal to 0"
            ]
            if (this.rules !== undefined) rules = rules.concat(this.rules)
            return rules
        }
    }
}
</script>
