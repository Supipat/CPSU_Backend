import logo from './logo.svg';
import './App.css';
import { BrowserRouter} from 'react-router-dom';
import AppRoutes from "./routes";
import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap/dist/js/bootstrap.bundle.min.js';

function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <AppRoutes />
      </BrowserRouter>
    </div>
  );
}

export default App;
