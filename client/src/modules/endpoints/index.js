import Table from "../../components/Table"

function EndpointsTable() {
  const columns = [ 
    {
        name: "rating",
        label: "Rating",
    },
    
    
    {
      name: "url",
      label: "URL",
      
      options: {
        filter: false,
        customBodyRender: (value, _tableMeta, _updateValue) => {
          return (
            <a href={value} target = "_blank" rel = "noopener noreferrer">{value}</a>
          );
        }
      }, 
      
    },
    
    {
      name: "length",
      label: "Length",
      
    },
    
    {
      name: "status",
      label: "Status",
      
    },
    
    {
      name: "redirectlocation",
      label: "Redirect Location",
      
      options: {
        filter: false,
        customBodyRender: (value, _tableMeta, _updateValue) => {
          return (
            <a href={value} target = "_blank" rel = "noopener noreferrer">{value}</a>
          );
        }
      }, 
      
    },
    
    
  ];

  return (
    <Table
      columns={columns}
      title="endpoints"
      api_url="http://127.0.0.1:8888/api/data"
    />
  )
}

export default {
  routeProps: {
      path: '/endpoints',
      component: EndpointsTable,
  },
  name: 'endpoints',
};
