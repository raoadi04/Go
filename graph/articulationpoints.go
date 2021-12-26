package graph

import "github.com/TheAlgorithms/Go/math/min"

type apHelper struct {
	is_ap              []bool
	visited            []bool
	child_cnt          []int
	discovery_time     []int
	earliest_discovery []int
}

// ArticulationPoint is a function to identify articulation points in a graph.
// The function takes the graph as an argument and returns a boolean slice which indicates whether a vertex is an articulation point or not.
// Worst Case Time Complexity: O(|V| + |E|)
// Auxiliary Space: O(|V|)
// reference: https://en.wikipedia.org/wiki/Biconnected_component and https://cptalks.quora.com/Cut-Vertex-Articulation-point
func ArticulationPoint(graph Graph) []bool {
	time := 0 // time variable to keep track of the time of discovery_time of a vertex

	//initialize all the variables
	apHelperInstance := &apHelper{
		is_ap:              make([]bool, graph.vertices),
		visited:            make([]bool, graph.vertices),
		child_cnt:          make([]int, graph.vertices),
		discovery_time:     make([]int, graph.vertices), // integer slice to store the discovery time of a vertex as we traverse the graph in a depth first manner
		earliest_discovery: make([]int, graph.vertices), // integer slice to store the earliest discovered vertex reachable from a vertex
	}
	articulationPointHelper(apHelperInstance, 0, -1, &time, graph)

	if apHelperInstance.child_cnt[0] == 1 {
		apHelperInstance.is_ap[0] = false // if the root has only one child, it is not an articulation point
	}

	return apHelperInstance.is_ap
}

// articulationPointHelper is a recursive function to traverse the graph and mark articulation points.
// Based on the depth first search transversal of the graph, however modified to keep track
// and update the child_cnt, discovery_time and earliest_discovery slices defined above
func articulationPointHelper(apHelperInstance *apHelper, vertex int, parent int, time *int, graph Graph) {
	apHelperInstance.visited[vertex] = true

	apHelperInstance.discovery_time[vertex] = *time                                       // Mark the time of discovery of a vertex
	apHelperInstance.earliest_discovery[vertex] = apHelperInstance.discovery_time[vertex] // set the earliest discovery time to the discovered time
	*time++                                                                               // increment the time

	for next_vertex := range graph.edges[vertex] {
		if next_vertex == parent {
			continue
		}

		if apHelperInstance.visited[next_vertex] { // if the vertex is already visited, then it is a back edge and manipulate earliest discovery time accordingly
			apHelperInstance.earliest_discovery[vertex] = min.Int(apHelperInstance.earliest_discovery[vertex], apHelperInstance.discovery_time[next_vertex])
		} else {
			apHelperInstance.child_cnt[vertex]++
			articulationPointHelper(apHelperInstance, next_vertex, vertex, time, graph)
			apHelperInstance.earliest_discovery[vertex] = min.Int(apHelperInstance.earliest_discovery[vertex], apHelperInstance.earliest_discovery[next_vertex])
			if apHelperInstance.earliest_discovery[next_vertex] >= apHelperInstance.discovery_time[vertex] {
				apHelperInstance.is_ap[vertex] = true
			}
		}
	}
}
