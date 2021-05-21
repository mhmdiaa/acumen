import modules from "./modules"
import { Route, Switch } from 'react-router-dom'

function Home() {
  return (
    <h1 align="center">Acumen</h1>
  )
}

function App() {

  let routes = modules.map(
    (m) => {
      return (
        <Route path={m.routeProps.path} component={m.routeProps.component}></Route>
      )
    }
  )

  return (
    <Switch>
      {routes}
      <Route component={Home} />
    </Switch>
  )
}

export default App;
