<template>
    <div>
        <v-row>
            <v-col cols="6" align-self="start">
                <h2>Analyzers</h2>
            </v-col>
            <v-col cols="6" align-self="end" class="text-right">
                <v-btn class="pa-2 ma-2" color="primary" @click="addAnalyzerHandler"> <v-icon>mdi-plus</v-icon> Install
                    Analyzer</v-btn>
            </v-col>
        </v-row>
        <v-row>
            <v-col cols="12">
                <InstallerModal :dialogVisible="visibleInstallerModal" @update:dialogVisible="visibleInstallerModal = $event" appType="analyzer"></InstallerModal>
                <v-table>
                    <thead>
                        <tr>
                            <th class="text-left">
                                Name
                            </th>
                            <th class="text-left">
                                Calories
                            </th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="item in desserts" :key="item.name">
                            <td>{{ item.name }}</td>
                            <td>{{ item.calories }}</td>
                        </tr>
                    </tbody>
                </v-table>
            </v-col>
        </v-row>
    </div>
</template>
<script>
import InstallerModal from '../components/InstallerModal.vue'
export default {
    components: {
        InstallerModal
    },
    methods: {
        addAnalyzerHandler() {
            this.visibleInstallerModal = true
        },
    },
    data() {
        return {
            visibleInstallerModal:false,
            desserts: [
                {
                    name: 'Frozen Yogurt',
                    calories: 159,
                },
                {
                    name: 'Ice cream sandwich',
                    calories: 237,
                },
                {
                    name: 'Eclair',
                    calories: 262,
                },
                {
                    name: 'Cupcake',
                    calories: 305,
                },
                {
                    name: 'Gingerbread',
                    calories: 356,
                },
                {
                    name: 'Jelly bean',
                    calories: 375,
                },
                {
                    name: 'Lollipop',
                    calories: 392,
                },
                {
                    name: 'Honeycomb',
                    calories: 408,
                },
                {
                    name: 'Donut',
                    calories: 452,
                },
                {
                    name: 'KitKat',
                    calories: 518,
                },
            ],
        }
    },
}
</script>