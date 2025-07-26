import { writeFileSync } from "fs";
import * as encoding from "lib0/encoding.js";
import { join } from "path";

// Generate test data for various data types using lib0 encoding
function generateTestData() {
  const testCases = [];

  const intCases = [
    0, 1, 127, 128, 255, 256, 16383, 16384, 65535, 65536, 1000000, -1, -127,
    -128, -1000,
  ];
  // Test Case 1: Various uint
  console.log("Generating uint test cases...");
  intCases
    .filter((x) => x > 0)
    .forEach((value, index) => {
      const encoder = encoding.createEncoder();
      encoding.writeVarUint(encoder, Math.abs(value));
      const data = encoding.toUint8Array(encoder);
      testCases.push({
        name: `uint_${index}_${value}`,
        type: "uint",
        value: value,
        encoded: Array.from(data),
      });
    });

  // Test Case 2: Various integers
  console.log("Generating uint test cases...");
  intCases.forEach((value, index) => {
    const encoder = encoding.createEncoder();
    encoding.writeVarInt(encoder, value);
    const data = encoding.toUint8Array(encoder);
    testCases.push({
      name: `int_${index}_${value}`,
      type: "int",
      value: value,
      encoded: Array.from(data),
    });
  });

  // Test Case 2: Strings
  console.log("Generating string test cases...");
  const stringCases = ["", "hello", "Hello, World!", "🌍", "A".repeat(300)];
  stringCases.forEach((value, index) => {
    const encoder = encoding.createEncoder();
    encoding.writeVarString(encoder, value);
    const data = encoding.toUint8Array(encoder);
    testCases.push({
      name: `string_${index}`,
      type: "string",
      value: value,
      encoded: Array.from(data),
    });
  });

  // Test Case 3: Uint8Arrays (binary data)
  console.log("Generating binary test cases...");
  const binaryCases = [
    new Uint8Array([]),
    new Uint8Array([1, 2, 3]),
    new Uint8Array([0, 255, 128]),
    new Uint8Array(Array.from({ length: 300 }, (_, i) => i % 256)),
  ];
  binaryCases.forEach((value, index) => {
    const encoder = encoding.createEncoder();
    encoding.writeVarUint8Array(encoder, value);
    const data = encoding.toUint8Array(encoder);
    testCases.push({
      name: `binary_${index}`,
      type: "binary",
      value: Array.from(value),
      encoded: Array.from(data),
    });
  });

  const floatCases = [
    0.0,
    1.0,
    -1.0,
    3.14159,
    1.23456789,
    Number.MAX_SAFE_INTEGER,
    Number.MIN_SAFE_INTEGER,
  ];
  // Test Case 4: Float32 numbers
  console.log("Generating float32 test cases...");
  floatCases.forEach((value, index) => {
    const encoder = encoding.createEncoder();
    encoding.writeFloat32(encoder, value);
    const data = encoding.toUint8Array(encoder);
    testCases.push({
      name: `float_${index}_${value.toString().replace(".", "_").replace("-", "neg")}`,
      type: "float32",
      value: value,
      encoded: Array.from(data),
    });
  });

  // Test Case 4: Float32 numbers
  console.log("Generating float64 test cases...");
  floatCases.forEach((value, index) => {
    const encoder = encoding.createEncoder();
    encoding.writeFloat64(encoder, value);
    const data = encoding.toUint8Array(encoder);
    testCases.push({
      name: `float_${index}_${value.toString().replace(".", "_").replace("-", "neg")}`,
      type: "float64",
      value: value,
      encoded: Array.from(data),
    });
  });

  const anyCases = [
    20,
    -20,
    3.14159,
    true,
    "hello world",
    ["hello world", true],
    { array: [], boolean: true },
  ];
  console.log("Generating any test cases...");
  anyCases.forEach((value, index) => {
    const encoder = encoding.createEncoder();
    encoding.writeAny(encoder, value);
    const data = encoding.toUint8Array(encoder);
    testCases.push({
      name: `any_${index}_${value.toString().replace(".", "_").replace("-", "neg")}`,
      type: "any",
      valuetype:
        typeof value === "number"
          ? Number.isInteger(value)
            ? "integer"
            : "float"
          : typeof value,
      value: value,
      encoded: Array.from(data),
    });
  });

  return testCases;
}

// Generate and save test data
const testData = generateTestData();
const output = {
  generated_at: new Date().toISOString(),
  lib0_version: "0.2.97",
  test_cases: testData,
};

writeFileSync(
  join(import.meta.dirname, "encoding_decoding_test_data.json"),
  JSON.stringify(output, null, 2),
);
console.log(`Generated ${testData.length} test cases`);
console.log("Test data saved to lib0_test_data.json");
