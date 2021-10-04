import React, { useState } from "react";
import axios from "axios";
import { Button, Container, Row, Col, Form, Alert } from "react-bootstrap";

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
      "^(https?:\\/\\/)?" + // protocol
        "((([a-z\\d]([a-z\\d-]*[a-z\\d])*)\\.)+[a-z]{2,}|" + // domain name
        "((\\d{1,3}\\.){3}\\d{1,3}))" + // OR ip (v4) address
        "(\\:\\d+)?(\\/[-a-z\\d%_.~+]*)*" + // port and path
        "(\\?[;&a-z\\d%_.~+=-]*)?" + // query string
        "(\\#[-a-z\\d_]*)?$", // fragment locator
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
        <Col lg="6" md="8" sm="12">
          <Form method="POST">
            <div className="text-center d-grid">
              <Form.Label className="m-3">
                <h1>URL Shortener</h1>
              </Form.Label>
              <Form.Control
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
              {shortUrl && (
                <Alert className="text-center" variant="success">
                  <a href={shortUrl}>{shortUrl}</a>
                </Alert>
              )}
              {error && (
                <Alert className="text-center" variant="danger">
                  URL is not valid
                </Alert>
              )}
            </div>
          </Form>
        </Col>
      </Row>
    </Container>
  );
};

export default URLShortener;
