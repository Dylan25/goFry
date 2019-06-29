import React from 'react';
import ReactDOM from 'react-dom';
import axios from 'axios';
import BlockCard from './BlockCard.js'
import Grid from '@material-ui/core/Grid';

class App extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            selectedFile: null
        }

    }

    fileSelectedHandler = event => {
        this.setState({
            selectedFile: event.target.files[0],
            fryedImage: {},
        })
    }

    fileUploadHandler = () => {
        console.log("clicked upload")
        const fd = new FormData();
        fd.append('image', this.state.selectedFile, this.state.selectedFile.name);
        axios.post('/api/', fd)
            .then(res => {
                console.log("hi");
                console.log(res);
                this.setState({
                    fryedImage: res.data.fryedImage
                })
            })
            .catch(error => {
                console.log("post error", error);
            })
    }

    createButtons = () => {
        return (
            <Grid
                container
                alignItems="center"
                justify="space-evenly"
            >
                <Grid item xs={1}>
                    <input type="file" onChange={this.fileSelectedHandler}/>
                </Grid>
                <Grid item xs={1}>
                    <button onClick={this.fileUploadHandler}>Upload</button>
                </Grid>
            </Grid>
        )
    }

    render() {
        return (
            <div>
                <BlockCard
                        payload={<img src={`data:image/jpeg;base64,${this.state.fryedImage}`}/>}
                        actions={
                            this.createButtons()
                        }
                />
            </div>
        );
    }
}

ReactDOM.render(
    <App />,
    document.querySelector('#root')
)