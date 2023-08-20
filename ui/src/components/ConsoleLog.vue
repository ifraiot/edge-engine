<template>
    <v-dialog v-model="modalVisible" max-width="300">
        <v-card>
            <v-card-title class="headline">Modal Title</v-card-title>
            <v-card-text>
                <div class="terminal-container">
                    <div class="terminal">
                        <pre ref="logContent" class="log-content">{{ logs }}</pre>
                    </div>
                </div>
            </v-card-text>
            <v-card-actions>
                <v-btn text @click="modalVisible = false">Close</v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>
<script>
// import axios from 'axios';
export default {
    props: {
        visible: {
            type: Boolean,
            default: false
        },
        logs: {
            type: String,
            default: ""
        }
    },
    data() {
        return {
            modalVisible: false
        }
    },
    methods: {
        closeModal() {
            this.$emit('update:modalVisible', false)
        }
    },
    watch: {
        logs: {
            immediate: true,
            handler(val) {
                this.log = this.logs +'\n'+ val
            }
            
        },
        visible: {
            immediate: true,
            handler(val) {
                this.modalVisible = val
            }
        }
    }
}

</script>