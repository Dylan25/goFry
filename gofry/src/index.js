import React from 'react';
import ReactDOM from 'react-dom';
import axios from 'axios';
import BlockCard from './BlockCard.js'
import Grid from '@material-ui/core/Grid';

class App extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            selectedFile: null,
            processingImage: false,
            fryedImage: null,
            selectedNewFile: false,
        }

    }

    fileSelectedHandler = event => {
        this.setState({
            selectedFile: event.target.files[0],
            fryedImage: event.target.files[0],
            selectedNewFile: true,

        })
    }

    fileUploadHandler = () => {
        const fd = new FormData();
        fd.append('image', this.state.selectedFile, this.state.selectedFile.name);
        this.setState({
            processingImage: true
        })
        axios.post('/api/', fd)
            .then(res => {
                this.setState({
                    fryedImage: res.data.fryedImage,
                    processingImage: false,
                    selectedNewFile: false,
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

    genImageJSX = () => {
        if(this.state.selectedNewFile) {
            return (<img src={URL.createObjectURL(this.state.selectedFile)}/>)
        } else {
            return (<img src={`data:image/jpeg;base64,${this.state.fryedImage}`}/>)
        }
    }

    render() {
        
        if(this.state.processingImage) {
            return (
                <div>
                    <BlockCard
                            payload={"Loading..."}
                            actions={
                                this.createButtons()
                            }
                    />
                </div>
            );
        } else if (this.state.selectedFile != null) {
            return (
                <div>
                    <BlockCard
                            payload={this.genImageJSX()}
                            actions={
                                this.createButtons()
                            }
                    />
                </div>
            );
        } else {
            return (
                <div>
                    <BlockCard
                            payload={" "}
                            actions={
                                this.createButtons()
                            }
                    />
                </div>
            );
        }
    }
}

ReactDOM.render(
    <App />,
    document.querySelector('#root')
)