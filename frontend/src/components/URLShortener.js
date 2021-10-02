import React, { useState, useEffect } from "react";
import axios from "axios";
import "./URLShortener.css";
import { Button, Container, Row, Col, Form } from "react-bootstrap";

const endpoint = "http://127.0.0.1:8080";

const URLShortener = () => {
  const [url, setUrl] = useState("");
  const [shortUrl, setShortUrl] = useState("");
  const [error, setErrorState] = useState(false);

  const handleURLChange = (e) => {
    setUrl(e.target.value);
  };

  const handleValidation = (url) => {
    let pattern = new RegExp(
      "^(https?:\\/\\/)?" +                                   // protocol
        "((([a-z\\d]([a-z\\d-]*[a-z\\d])*)\\.)+[a-z]{2,}|" +  // domain name
        "((\\d{1,3}\\.){3}\\d{1,3}))" +                       // OR ip (v4) address
        "(\\:\\d+)?(\\/[-a-z\\d%_.~+]*)*" +                   // port and path
        "(\\?[;&a-z\\d%_.~+=-]*)?" +                          // query string
        "(\\#[-a-z\\d_]*)?$",                                 // fragment locator
      "i"
    ); 
    const regex = new RegExp(pattern);

    return url.match(regex);
  };

  async function handleButtonOnClick(e) {
    e.preventDefault();
    setShortUrl("");
    if (handleValidation(url)) {
      const response = await axios.post(
        endpoint + "/URL",
        url.split(":")[0] !== "http" && url.split(":")[0] !== "https"
          ? "https://" + url
          : url,
        {
          headers: {
            "Content-Type": "application/x-www-form-urlencoded",
          },
        }
      );
      setErrorState(false);
      setShortUrl(response.data);
    } else {
      setErrorState(true);
    }
  }

  return (
    <Container>
      <Row className="justify-content-md-center">
        <Col xs lg="3">
          <Form method="POST">
            <div className="text-center">
              <Form.Label className="m-3">URL Shortener</Form.Label>
              <Form.Control
                name="url"
                type="url"
                placeholder="Your URL"
                onChange={handleURLChange}
              />
            </div>
            <div className="d-grid gap-2">
              <Button
                type="submit"
                variant="primary"
                size="lg"
                onClick={handleButtonOnClick}
              >
                Shorten Your URL
              </Button>
              {shortUrl && <a href={shortUrl}>{shortUrl}</a>}
              {error && <div>URL is not valid </div>}
            </div>
          </Form>
        </Col>
      </Row>
    </Container>
  );
};

export default URLShortener;
