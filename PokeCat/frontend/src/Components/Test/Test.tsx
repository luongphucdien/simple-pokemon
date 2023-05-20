import { Col, Form, Input, Row } from "antd"
import { useRef } from "react"

export const Test = () => {
    interface TestData {
        test: string
    }

    const handleSend = (testData: TestData) => {

    }

    return (
        <>
            <Row justify={"center"}>
                <Col span={8}>
                    <Form
                        name="test-data"
                        onFinish={handleSend}
                        labelCol={{span: 8}}
                    >
                        <Form.Item
                            label={"Test"}
                            name={"test"}
                        >
                            <Input/>
                        </Form.Item>
                    </Form>
                </Col>
            </Row>
        </>
    )
}