import { Pie, mixins } from 'vue-chartjs'
const { reactiveProp } = mixins

export default {
    mixins: [Pie,reactiveProp],
    props: ['options'],
    mounted() {
        this.renderChart(this.data, this.options)
    }
}