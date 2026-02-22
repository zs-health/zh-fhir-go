# 🧑‍💻 NO-CODER-GUIDE: Understanding and Using the Bangladesh Core FHIR IG

Welcome! This guide is designed for anyone who wants to understand and use the Bangladesh Core FHIR Implementation Guide (IG) without needing technical coding knowledge. Here, we'll explain what everything means and how to find the information you need.

## What is FHIR and an Implementation Guide (IG)?

Imagine healthcare systems speaking different languages. **FHIR (Fast Healthcare Interoperability Resources)** is like a universal translator, allowing these systems to understand each other. It defines standard ways to represent and exchange healthcare information, such as patient records, medications, and lab results.

An **Implementation Guide (IG)** is a detailed instruction manual that tells you exactly how to use FHIR for a specific purpose or in a particular region. Our Bangladesh Core FHIR IG provides the rules and definitions for exchanging health data within Bangladesh.

## 🌐 How to View the Published FHIR IG

Your Bangladesh Core FHIR IG is published online and automatically updated. You can access it here:

*   **Bangladesh Core FHIR IG**: [https://zs-health.github.io/zh-fhir-go/](https://zs-health.github.io/zh-fhir-go/)

When you open this link, you'll see a website that looks like a comprehensive document. This is your Implementation Guide.

## 🗺️ Understanding the Structure of the IG

The IG is organized into several sections to help you navigate. Here are some key areas you'll find:

*   **Home/Introduction**: Provides an overview of the IG, its purpose, and what you can expect to find.
*   **Profiles**: This is a very important section! **Profiles** define how standard FHIR resources (like a Patient record or an Encounter) are adapted for use in Bangladesh. For example, the `BDPatient` profile will specify how a patient's information, including national ID, is represented in Bangladesh.
    *   **How to Find Specific Profiles**: Look for a navigation link usually titled "Profiles" or "Artifacts". Clicking on it will show a list of all defined profiles, such as `BDPatient`, `BDAddress`, `BDEncounter`, etc.
*   **Terminology**: This section explains the **Code Systems** and **Value Sets** used in the IG. These are standardized lists of codes (like ICD-11 for diagnoses or specific codes for administrative divisions) that ensure everyone uses the same terms.
    *   **How to Read Terminology**: Navigate to the "Terminology" or "Code Systems" section. Here you'll find definitions and lists of codes used throughout the IG.
*   **API Documentation**: If you're interested in how systems interact with the FHIR server, this section provides details on the available API endpoints and how to use them.

## ❓ Glossary of Terms

To help you understand the jargon, here's a quick glossary:

*   **FHIR (Fast Healthcare Interoperability Resources)**: A standard for exchanging healthcare information electronically.
*   **Implementation Guide (IG)**: A document that provides specific rules and guidance for implementing FHIR in a particular context.
*   **FSH (FHIR Shorthand)**: A simplified language used by developers to define FHIR profiles and extensions.
*   **Profile**: A customization of a standard FHIR resource to meet specific local or use-case requirements (e.g., `BDPatient` is a profile of the standard FHIR `Patient` resource).
*   **Extension**: A way to add new data elements to FHIR resources that are not part of the standard specification.
*   **Code System**: A comprehensive list of codes and their meanings (e.g., ICD-11 for diseases).
*   **Value Set**: A subset of codes from one or more code systems, used for a specific purpose (e.g., a value set for types of blood groups).

## 🤝 How to Report Issues or Suggest Improvements

Your feedback is crucial for improving this Implementation Guide. If you find any errors, have suggestions for new content, or notice anything unclear, please:

1.  Go to the [GitHub Issues page](https://github.com/zs-health/zh-fhir-go/issues).
2.  Click on the "New issue" button.
3.  Provide a clear title and description of the issue or suggestion.

Thank you for helping us improve healthcare interoperability in Bangladesh!
