<template>
    <div style="height: 100vh">
        <VueFlow v-model="elements" :default-edge-options="{ type: 'smoothstep' }" class="basicflow" />

    </div>
</template>
<script>
import { VueFlow } from '@vue-flow/core'
import { Position } from '@vue-flow/core'
import axios from 'axios';
export default {
    components: {
        VueFlow
    },
    async created() {
        this.installedApplications = await this.fetchInstalledApplicationsAPI()
        this.connectors = this.installedApplications.filter(item => item.application.type === 'connector')
        this.analyzers = this.installedApplications.filter(item => item.application.type === 'analyzer')
        this.integrators = this.installedApplications.filter(item => item.application.type === 'integrator')
        console.log(this.installedApplications)

        //connectors
        let y = 200
        this.connectors.forEach(item => {
            //Node
            this.elements.push(
                { id: `node_connector_${item.id}`, type: 'input', label: item.application.label, position: { x: 100, y: y }, sourcePosition: Position.Right }
            )

            //Edge
            this.elements.push(
                { id: `edge_connector_${item.id}_to_bus`, source: `node_connector_${item.id}`, target: 'bus', animated: true }
            )
            y += 100
        })

        //integrators
        y = 200
        this.integrators.forEach(item => {
            //Node
            this.elements.push(
                {
                    id: `node_integrator_${item.id}`,
                    type: 'output',
                    label: item.application.label,
                    position: { x: 900, y: y },
                    sourcePosition: Position.Right,
                    targetPosition: Position.Left,
                }
            )

            //Edge
            this.elements.push(
                { id: `edge_connector_${item.id}_to_bus`, source: 'bus', target: `node_integrator_${item.id}`, animated: true }
            )
            y += 100
        })


        //analyzers
        // y = 500
        // this.analyzers.forEach(item => {
        //     //Node
        //     this.elements.push(
        //         { id: `node_analyzer_${item.id}`, type: 'input', label: item.application.label, position: { x: 500, y: y }, sourcePosition: Position.Left }
        //     )

        //     //Edge
        //     this.elements.push(
        //         { id: `node_analyzer_${item.id}_to_bus_in`, source: 'bus', target: `node_analyzer_${item.id}`, animated: true }
        //     )
        //     this.elements.push(
        //         { id: `node_analyzer_${item.id}_to_bus_out`, target: 'bus', source: `node_analyzer_${item.id}`, animated: true }
        //     )
        //     y += 100
        // })
    },
    data() {
        return {
            connectors: [],
            analyzers: [],
            integrators: [],
            installedApplications: [],
            elements: [
                { id: 'bus', label: 'Message Queue', position: { x: 500, y: 200 }, sourcePosition: Position.Right, targetPosition: Position.Left, },
            ]
            // elements: [
            //     // Nodes
            //     // An input node, specified by using `type: 'input'`
            //     { id: 'connector-1', type: 'input', label: 'Node 1', position: { x: 100, y: 200 }, sourcePosition: Position.Right },

            //     // Default nodes, you can omit `type: 'default'`


            //     // An output node, specified by using `type: 'output'`
            //     { id: 'integrator-1', type: 'output', label: 'Node 4', position: { x: 900, y: 200 }, targetPosition: Position.Left, },

            //     // Edges
            //     // An animated edge
            //     { id: 'e1-2', source: 'connector-1', target: 'bus', animated: true },
            //     { id: 'e2-4', source: 'bus', target: 'integrator-1', animated: true },
            // ]
        };
    },
    methods: {
        async fetchInstalledApplicationsAPI() {
            try {
                var installedApplications = []
                const response = await axios.get('http://localhost:8000/api/applications');
                const items = response.data.data;
                for (let i = 0; i < items.length; i++) {
                    items[i].config = JSON.parse(items[i].config);
                    installedApplications.push(items[i]);
                }
                return installedApplications
            } catch (error) {
                console.error('Error fetching items:', error);
            }
            return []
        },
    },
};
</script>
<style></style>