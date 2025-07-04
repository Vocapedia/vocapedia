import { ref } from 'vue';
import { jsPDF } from "jspdf";
import { i18n } from "@/i18n/i18n";
import { getLangByCode } from "@/utils/language/languages";
import "@/assets/Roboto-bold";
import "@/assets/Roboto-normal";

export function usePDFGenerator() {
    const isPDFGenerating = ref(false);
    
    let savedStyles = {
        links: [],
        styles: [],
        inlineStyles: []
    };

    const saveAndRemoveStyles = () => {
        savedStyles.links = Array.from(document.querySelectorAll('link[rel="stylesheet"]'));
        savedStyles.styles = Array.from(document.querySelectorAll('style'));
        savedStyles.inlineStyles = Array.from(document.querySelectorAll('*')).map(element => ({
            element,
            style: element.getAttribute('style')
        }));
        savedStyles.links.forEach(link => link.remove());
        savedStyles.styles.forEach(style => style.remove());
        savedStyles.inlineStyles.forEach(item => item.element.removeAttribute('style'));
    };

    const restoreStyles = () => {
        savedStyles.links.forEach(link => document.head.appendChild(link));
        savedStyles.styles.forEach(style => document.head.appendChild(style));
        savedStyles.inlineStyles.forEach(item => {
            if (item.style) {
                item.element.setAttribute('style', item.style);
            }
        });
    };

    const generatePDF = async (chapterData) => {
        isPDFGenerating.value = true;
        const doc = new jsPDF();
        const marginX = 14;
        const startY = 20;

        saveAndRemoveStyles();

        const tutorialDiv = document.createElement('div');

        // Title
        const titleEl = document.createElement('h1');
        titleEl.textContent = chapterData.title;
        titleEl.style.fontSize = '20px';
        titleEl.style.fontWeight = 'bold';
        titleEl.style.marginBottom = '8px';
        titleEl.style.fontFamily = "'Roboto', sans-serif";
        tutorialDiv.appendChild(titleEl);

        // Description
        const descEl = document.createElement('p');
        descEl.textContent = chapterData.description;
        descEl.style.marginBottom = '6px';
        descEl.style.fontFamily = "'Roboto', sans-serif";
        tutorialDiv.appendChild(descEl);

        // Creator
        const userEl = document.createElement('p');
        userEl.textContent = `@${chapterData.creator.username}`;
        userEl.style.color = '#666';
        userEl.style.marginBottom = '14px';
        userEl.style.fontFamily = "'Roboto', sans-serif";
        tutorialDiv.appendChild(userEl);

        tutorialDiv.style.width = '160mm';

        // Tutorial content
        const innerDiv = document.createElement('div');
        innerDiv.style.width = '100%';
        innerDiv.style.fontFamily = "'Roboto', sans-serif";
        innerDiv.innerHTML = chapterData.tutorial;

        // Process images
        const images = innerDiv.querySelectorAll('img');
        innerDiv.querySelectorAll('iframe').forEach(iframe => {
            iframe.remove();
        });
        images.forEach(img => {
            img.style.maxWidth = '50%';
            img.style.width = '100%';
            img.style.height = 'auto';
            img.style.display = 'block';
            img.style.marginLeft = 'auto';
            img.style.marginRight = 'auto';
        });

        tutorialDiv.appendChild(innerDiv);

        // Words table
        const tableContainer = document.createElement('div');
        tableContainer.style.fontFamily = "'Roboto', sans-serif";
        tableContainer.style.gap = '16px';
        tableContainer.style.width = '100%';
        
        const wb = chapterData.word_bases;
        tableContainer.innerHTML = `
            <table style="width: 100%; border-collapse: separate;border-spacing: 0 8px; font-size: 13px; font-family: 'Roboto', sans-serif;">
              <thead>
                <tr style="background-color: #2980b9; color: #fff; font-size: 15px;">
                  <th style="border: 1px solid #2980b9; padding: 15px 12px; border-top-left-radius: 6px;">${i18n.global.t("chapter.type")}</th>
                  <th style="border: 1px solid #2980b9; padding: 15px 12px;">${i18n.global.t("chapter.language")}</th>
                  <th style="border: 1px solid #2980b9; padding: 15px 12px; border-top-right-radius: 6px;">${i18n.global.t("chapter.word")}</th>
                </tr>
              </thead>
              <tbody>
                ${wb.map((base) =>
                    base.words.map((word, i) => `
                      ${i % 2 == 0 ? '<tr style="background-color: #ecf0f1">' : '<tr>'}
                        ${i === 0 ? `
                          <td rowspan="${base.words.length}" style="border-top: 0.5px solid #ccc;border-bottom: 0.5px solid #ccc;border-left: 0.5px solid #ccc; padding: 6px 8px; text-align:center; font-weight: 500; background-color: #ecf0f1;">
                            ${i18n.global.t("word_types." + base.type)}
                          </td>` : ``}
                          <td style="padding: 10px 8px; text-align:center;">
                            ${getLangByCode(word.lang).name}
                          </td>
                          <td style="padding: 10px 8px;">
                            ${word.word}
                          </td>
                        </tr>
                      `).join('')
                ).join('')}
              </tbody>
            </table>
        `;

        tutorialDiv.appendChild(tableContainer);
        document.body.appendChild(tutorialDiv);

        await doc.html(tutorialDiv, {
            useCORS: true,
            foreignObjectRendering: true,
            x: marginX,
            y: startY,
            html2canvas: { scale: 0.3 },
            callback: () => {
                restoreStyles();
                document.body.removeChild(tutorialDiv);
            }
        });

        doc.save(`${chapterData.title}.pdf`);
        isPDFGenerating.value = false;
    };

    return {
        isPDFGenerating,
        generatePDF
    };
}
