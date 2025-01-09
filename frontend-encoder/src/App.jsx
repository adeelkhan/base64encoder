import { useState } from "react";
import "./App.css";

function App() {
  const [plainTextEdit, setPlainTextEdit] = useState("");
  const [encodedString, setEncodedString] = useState("");
  const [decodedStringEdit, setDecodedStringEdit] = useState("");
  const [decodedString, setDecodedString] = useState("");

  const editPlainText = (e) => {
    setPlainTextEdit(e.target.value);
  };

  const editDecodedText = (e) => {
    setDecodedStringEdit(e.target.value);
  };

  const base64EncodedHandler = () => {
    const options = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        PlainText: plainTextEdit,
      }),
    };
    fetch("http://localhost:8000/encode", options)
      .then((response) => response.json())
      .then((data) => {
        setEncodedString(data["Base64EncodedString"]);
      })
      .catch((error) => console.log(error));
  };
  const base64DecodeddHandler = () => {
    console.log("decoder called");
    const options = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        Base64EncodedString: decodedStringEdit,
      }),
    };
    fetch("http://localhost:8000/decode", options)
      .then((response) => response.json())
      .then((data) => {
        setDecodedString(data["PlainText"]);
      })
      .catch((error) => console.log(error));
  };

  return (
    <>
      <h1>Base64 Encoder/Decoder Tool</h1>
      <div className="card">
        <div>
          <p>String to Encode:</p>
          <div>
            <div>
              <textarea
                value={plainTextEdit}
                name="string_to_encode"
                rows="4"
                cols="50"
                onChange={(e) => editPlainText(e)}
              ></textarea>
              <br />
              <button onClick={base64EncodedHandler}>Encode</button>
            </div>
            <div>
              <p>Encoded String:</p>
              <textarea
                value={encodedString}
                type="textarea"
                name="encoded_string"
                rows="4"
                cols="50"
              ></textarea>
            </div>
          </div>
        </div>
        <br />
        <div>
          <p>String to Decode:</p>
          <div>
            <div>
              <textarea
                value={decodedStringEdit}
                type="textarea"
                name="string_to_decode"
                rows="4"
                cols="50"
                onChange={(e) => editDecodedText(e)}
              ></textarea>
              <br />
              <button onClick={base64DecodeddHandler}>Decode</button>
            </div>
            <div>
              <p>Decoded String:</p>
              <textarea
                value={decodedString}
                type="textarea"
                name="decoded_string"
                rows="4"
                cols="50"
              ></textarea>
            </div>
          </div>
        </div>
      </div>
    </>
  );
}

export default App;
